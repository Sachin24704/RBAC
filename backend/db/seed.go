package db

import (
	"backend/ent"
	"context"
	"log"
)

func SeedDatabase(client *ent.Client) {
	// Check if roles already exist to avoid reseeding
	ctx := context.Background()
	count, err := client.Role.Query().Count(ctx)
	if err != nil {
		log.Fatalf("failed to check roles: %v", err)
	}
	if count > 0 {
		log.Println("Seed data already exists, skipping seeding.")
		return
	}

	log.Println("Seeding initial data...")

	// Create Roles
	adminRole, _ := client.Role.Create().SetName("Admin").Save(ctx)
	managerRole, _ := client.Role.Create().SetName("Manager").Save(ctx)

	// Create Users
	alice, _ := client.User.Create().SetName("Ram").Save(ctx)
	bob, _ := client.User.Create().SetName("Sham").Save(ctx)

	// Assign Roles to Users
	client.UserRole.Create().SetUser(alice).SetRole(adminRole).Save(ctx)
	client.UserRole.Create().SetUser(bob).SetRole(managerRole).Save(ctx)

	// Create Models
	leaveModel, _ := client.Model.Create().SetName("Leave").Save(ctx)
	projectModel, _ := client.Model.Create().SetName("Project").Save(ctx)

	// Create Operations for Leave model
	approveOp, _ := client.Operation.Create().SetName("Approve").SetModel(leaveModel).Save(ctx)
	rejectOp, _ := client.Operation.Create().SetName("Reject").SetModel(leaveModel).Save(ctx)
	forwardOp, _ := client.Operation.Create().SetName("Forward").SetModel(leaveModel).Save(ctx)

	// Create Operations for Project model
	viewOp, _ := client.Operation.Create().SetName("View").SetModel(projectModel).Save(ctx)
	editOp, _ := client.Operation.Create().SetName("Edit").SetModel(projectModel).Save(ctx)
	commentOp, _ := client.Operation.Create().SetName("Comment").SetModel(projectModel).Save(ctx)

	// Assign Operations to Roles
	client.RoleOperation.Create().SetRole(adminRole).SetOperations(approveOp).Save(ctx)
	client.RoleOperation.Create().SetRole(adminRole).SetOperations(rejectOp).Save(ctx)
	client.RoleOperation.Create().SetRole(adminRole).SetOperations(forwardOp).Save(ctx)

	client.RoleOperation.Create().SetRole(managerRole).SetOperations(viewOp).Save(ctx)
	client.RoleOperation.Create().SetRole(managerRole).SetOperations(editOp).Save(ctx)
	client.RoleOperation.Create().SetRole(managerRole).SetOperations(commentOp).Save(ctx)

	log.Println("Seeding completed successfully!")
}
