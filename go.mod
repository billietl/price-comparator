module gitlab.com/billietl/price-comparator

go 1.16

replace gitlab.com/billietl/price-comparator/store => ./store

replace gitlab.com/billietl/price-comparator/firestoreclient => ./firestoreclient

require (
	cloud.google.com/go/firestore v1.5.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	gitlab.com/billietl/price-comparator/firestoreclient v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/billietl/price-comparator/store v0.0.0-00010101000000-000000000000 // indirect
)
