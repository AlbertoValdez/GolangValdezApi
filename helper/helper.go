package helper

//Catch manejador de errores
func Catch(err error) {

	if err != nil {
		panic(err)
	}

}
