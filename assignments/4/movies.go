package main
import "fmt"
type movies struct {
	name string
language string
 releaseDate string
 director string
  producer string 
   duration int
}
func main() {
	var m=map[string]movies {"MA": movies{"Spideman","English","10/10/2015","Speilsberg","Richie Rich",120},
	"MB": movies{"Batman","English","10/10/2016","Speilsberg","Richie Rich",120},
	"MC": movies{"Ironman","English","10/10/2017","Speilsberg","Richie Rich",120},
	"MD": movies{"Spideman","English","10/10/2018","Speilsberg","Richie Rich",120}}
for _,v:=range m {
	fmt.Println(v.name,v.language,v.releaseDate,v.director,v.producer,v.duration,"\n")

}
var movname string
fmt.Println("Enter a movie name to delete")
fmt.Scanln(&movname)
delete(m,movname)
for _,v:=range m {
	fmt.Println(v.name,v.language,v.releaseDate,v.director,v.producer,v.duration,"\n")

}
}