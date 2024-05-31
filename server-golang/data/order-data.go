package data

import "grpc-go/main/proto"

var ORDERS = []proto.OrderResponse{
	{
		Date: 100,
		From: &proto.Address{Name: "Hans", City: "Leipzig"},
		Items: []*proto.OrderItem{
			{Title: "Copy paper", Price: 25},
			{Title: "Laser printer", Price: 250},
			{Title: "Cat-7 cable", Price: 20},
		},
	},
	{
		Date: 200,
		From: &proto.Address{Name: "Inge", City: "Berlin"},
		Items: []*proto.OrderItem{
			{Title: "Screwdriver", Price: 10},
			{Title: "Wrench", Price: 5},
		},
	},
	{
		Date: 300,
		From: &proto.Address{Name: "Fred", City: "Munich"},
		Items: []*proto.OrderItem{
			{Title: "Filter cone", Price: 15},
		},
	},
	{
		Date: 400,
		From: &proto.Address{Name: "Anna", City: "Zurich"},
		Items: []*proto.OrderItem{
			{Title: "Laptop", Price: 1500},
			{Title: "Power supply", Price: 15},
		},
	},
	{
		Date: 500,
		From: &proto.Address{Name: "Kurt", City: "Vienna"},
		Items: []*proto.OrderItem{
			{Title: "Stopwatch", Price: 120},
		},
	},
}
