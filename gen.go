package random

//go:generate mockery -name Interface -inpkg -case=underscore
//go:generate sed -i s/MockInterface/Mock/ mock_interface.go
//go:generate sed -i s/_m/m/ mock_interface.go
//go:generate mv mock_interface.go mock.go
