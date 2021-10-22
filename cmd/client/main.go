package main

func main(){
	flag.Parse();
	if flag.NArg() < 2 {
		log.Fatal()
	}
}