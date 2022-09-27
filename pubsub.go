package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PubSubTest() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "svc-project-1fae")
	if err != nil {
		fmt.Println("err-NewClient :", err)
		// TODO: Handle error.
	}

	topic := client.Topic("mediahub-job-notifications")

	// Create a new topic with the given name.
	// topic, err := client.CreateTopic(ctx, "topicName")
	fmt.Println("topic :", topic)
	if err != nil {
		fmt.Println("err-CreateTopic :", err)
		// TODO: Handle error.
	}

	// Create a new subscription to the previously created topic
	// with the given name.
	// sub, err := client.CreateSubscription(ctx, "subName", pubsub.SubscriptionConfig{
	// 	Topic:            topic,
	// 	AckDeadline:      10 * time.Second,
	// 	ExpirationPolicy: 25 * time.Hour,
	// })
	sub := client.Subscription("subName")
	// if err != nil {
	// 	fmt.Println("err-CreateSubscription :", err)
	// 	// TODO: Handle error.
	// }

	msg := &pubsub.Message{
		Data: []byte("Hello, world!"),
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		fmt.Println("Publish-err :", err)
		return
	}

	fmt.Println("Message published")

	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("msg Received :", string(msg.Data))
		msg.Ack()
	})
	_ = sub // TODO: use the subscription.
}
