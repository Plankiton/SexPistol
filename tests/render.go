package main
import "github.com/Plankiton/SexPistol/Html"

func main() {
    j := map[string]interface{} {
        "name": "joao",
        "old": 34,
    }

    res, _ := SexHtml.Render(`
<div>
{{.name}} - {{.old}}
</div>
    `, j)
    print(string(res), "\n")
}
