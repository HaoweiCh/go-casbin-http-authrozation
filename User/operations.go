package User

var (
	Authorized Items
)

func init() {
	Authorized = append(Authorized, Item{
		ID:   1,
		Name: "Admin",
		Role: "admin",
	})
	Authorized = append(Authorized, Item{
		ID:   2,
		Name: "Sabine",
		Role: "member",
	})
	Authorized = append(Authorized, Item{
		ID:   3,
		Name: "Sepp",
		Role: "member",
	})
}
