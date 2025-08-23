package handlers

import (
	"github.com/AlperSeyman/hrms-fiber-mongodb/database"
	"github.com/AlperSeyman/hrms-fiber-mongodb/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllEmployees(c *fiber.Ctx) error {
	collection := database.GetCollection()
	filter := bson.D{{}}

	cur, err := collection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer cur.Close(c.Context())

	var employees []models.Employee
	if err := cur.All(c.Context(), &employees); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {

	collection := database.GetCollection()

	params := c.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	filter := bson.M{"_id": _id}

	var employee models.Employee

	if err := collection.FindOne(c.Context(), filter).Decode(&employee); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employee)

}

func CreateEmployee(c *fiber.Ctx) error {

	collection := database.GetCollection()

	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	employee.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(employee)

}

func UpdateEmployee(c *fiber.Ctx) error {

	collection := database.GetCollection()

	params := c.Params("id")

	var employee models.Employee

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	filter := bson.M{"_id": _id}

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	update := bson.M{
		"$set": bson.M{
			"name":   employee.Name,
			"salary": employee.Salary,
			"age":    employee.Age,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedEmployee models.Employee

	err = collection.FindOneAndUpdate(c.Context(), filter, update, opts).Decode(&updatedEmployee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedEmployee)

}

func DeleteEmployee(c *fiber.Ctx) error {

	collection := database.GetCollection()

	params := c.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": _id}

	collection.DeleteOne(c.Context(), filter)

	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}
