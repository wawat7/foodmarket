package role

const Admin = "admin"
const Customer = "customer"

func List() map[string]string {
	return map[string]string{
		Admin:    "Administrator",
		Customer: "Customer",
	}
}
