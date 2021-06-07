package main
import (
    "github.com/Plankiton/SexPistol/Casing"
    "os"
)

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

    file, _ := os.Open("render.html")
    defer file.Close()

    res, _ = SexHtml.Render(file, j)
    print(string(res), "\n")
}
