package handlers

import (
    "better/database"
    "better/models"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get all messages.
// @Description Fetch every message available.
// @Tags Messages
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Message
// @Router /messages [get]
func HandleAllMessages(c *fiber.Ctx) error {
    coll := database.GetCollection("messages")

    filter := bson.M{}
    opts := options.Find().SetSkip(0).SetLimit(100)

    cursor, err := coll.Find(c.Context(), filter, opts)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }

    messages := make([]models.Message, 0)
    if err = cursor.All(c.Context(), &messages); err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }

    return c.Status(200).JSON(messages)
}

type CreateMessage struct {
    Username    string             `json:"username" bson:"username"`
    Content     string             `json:"content" bson:"content"`
    Channel     string             `json:"channel" bson:"channel"`
    Date        string             `json:"date" bson:"date"`
}

type CreateMessageRes struct {
    InsertedId primitive.ObjectID `json:"inserted_id" bson:"_id"`
}

// @Summary Get all messages in a channel.
// @Description Fetch every message in a channel.
// @Tags Messages
// @Param channel path string true "Channel ID"
// @Produce json
// @Success 200 {object} models.Message
// @Router /messages/:channel [get]
func HandleGetMessagesInChannel(c *fiber.Ctx) error {
    channel := c.Params("channel")

    dbId, err := primitive.ObjectIDFromHex(channel)
    /*if err != nil {
    return c.Status(400).JSON(fiber.Map{"invalid channel": err.Error()})
    }*/

    fmt.Println(dbId)

    coll := database.GetCollection("messages")
    filter := bson.M{"channel": channel}
    //var message models.Message

    opts := options.Find().SetSkip(0).SetLimit(100)


    cursor, err := coll.Find(c.Context(), filter, opts)
    messages := make([]models.Message, 0)
    if err = cursor.All(c.Context(), &messages); err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }


    return c.Status(200).JSON(messages)
}

// @Summary Create a message.
// @Description Create a single message.
// @Tags Messages
// @Accept json
// @Param message body CreateMessage true "Message to create"
// @Produce json
// @Success 200 {object} CreateMessageRes
// @Router /messages [post]
func HandleCreateMessage(c *fiber.Ctx) error {
    newMessage := new(CreateMessage)

    if err := c.BodyParser(newMessage); err != nil {
        return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
    }

    coll := database.GetCollection("messages")
    res, err := coll.InsertOne(c.Context(), newMessage)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"inserted_id": res.InsertedID})
}

type UpdateMessage struct {
    Username    string             `json:"username" bson:"username"`
    Content     string             `json:"content" bson:"content"`
    Channel     string             `json:"channel" bson:"channel"`
    Date        string             `json:"date" bson:"date"`
}

type UpdateMessageRes struct {
    UpdatedCount int64 `json:"updated_count" bson:"updated_count"`
}

// @Summary Update a message.
// @Description Update a single message.
// @Tags Messages
// @Accept json
// @Param todo body UpdateMessage true "Message update data"
// @Param id path string true "Message ID"
// @Produce json
// @Success 200 {object} UpdateMessageRes
// @Router /messages/:id [put]
func HandleUpdateMessage(c *fiber.Ctx) error {
    id := c.Params("id")
    dbId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
    }

    updatedMessage := new(UpdateMessage)

    if err := c.BodyParser(updatedMessage); err != nil {
        return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
    }

    coll := database.GetCollection("messages")
    filter := bson.M{"_id": dbId}
    update := bson.M{"$set": updatedMessage}
    res, err := coll.UpdateOne(c.Context(), filter, update)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"updated_count": res.ModifiedCount})
}

/*
type DeleteTodoResDTO struct {
    DeletedCount int64 `json:"deleted_count" bson:"deleted_count"`
}
*/

/*
// @Summary Delete a single message.
// @Description Delete a single message by id.
// @Tags messages
// @Param id path string true "Message ID"
// @Produce json
// @Success 200 {object} DeleteTodoResDTO
// @Router /todos/:id [delete]
func HandleDeleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    dbId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
    }

    coll := database.GetCollection("messages")
    filter := bson.M{"_id": dbId}
    res, err := coll.DeleteOne(c.Context(), filter)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"deleted_count": res.DeletedCount})
}
*/