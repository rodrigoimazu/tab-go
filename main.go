package main

import (
    "embed"
    "fmt"
    "net/http"
    "os"
    "strconv"
)

//go:embed style.css
var staticFiles embed.FS

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
    http.HandleFunc("/", formHandler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	fmt.Printf("🚀 http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    numStr := r.FormValue("num")
    htmlTabuada := ""

    if r.Method == http.MethodPost && numStr != "" {
        num, _ := strconv.Atoi(numStr)
        htmlTabuada = `
        <div class="tabuada">
            <h1>Tabuada do <span>` + numStr + `</span></h1>
            ` + linhaTabuada(num) + `
			<br>
			<div>
                <form action="/" method="GET">
                    <button type="submit">Limpar</button>
                </form>
            </div>
        </div>`
    }

    html := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="/static/style.css">
    <title>Tabuada Go</title>
</head>
<body>
    <h1>Tabuada</h1>
    <form action="/" method="POST">
        <label for="number">Digite um número (1-99): </label>
        <input name="num" type="number" min="1" max="99" required id="number" autocomplete="off">
        <button type="submit">Calcular</button>
    </form>
    ` + htmlTabuada + `
</body>
</html>`
    fmt.Fprint(w, html)
}

func linhaTabuada(num int) string {
    html := ""
    for i := 1; i <= 10; i++ {
        html += fmt.Sprintf(`<p><b>%d</b> × <b>%d</b> = %d</p>`,
            num, i, num*i)
    }
    return html
}
