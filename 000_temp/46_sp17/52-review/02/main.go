package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<style>
    html, body, h1 {
        padding: 0;
        border: 0;
        margin: 0;
        box-sizing: border-box;
    }

    main {
        display: flex;
        justify-content: center;
        align-items: center;
        background-image: linear-gradient(red, yellow, blue);
        height: 100vh;
    }

    h1 {
        font-size: 8rem;
        color: white;
    }

</style>
<body>

<main>
<h1>hello world</h1>
</main>

<ul>
    <li>001</li>
    <li>002</li>
    <li>003</li>
    <li>004</li>
    <li>005</li>
    <li>006</li>
    <li>007</li>
    <li>008</li>
    <li>009</li>
    <li>010</li>
    <li>011</li>
    <li>012</li>
    <li>013</li>
    <li>014</li>
    <li>015</li>
    <li>016</li>
    <li>017</li>
    <li>018</li>
    <li>019</li>
    <li>020</li>
    <li>021</li>
    <li>022</li>
    <li>023</li>
    <li>024</li>
    <li>025</li>
    <li>026</li>
    <li>027</li>
    <li>028</li>
    <li>029</li>
    <li>030</li>
    <li>031</li>
    <li>032</li>
    <li>033</li>
    <li>034</li>
    <li>035</li>
    <li>036</li>
    <li>037</li>
    <li>038</li>
    <li>039</li>
    <li>040</li>
    <li>041</li>
    <li>042</li>
    <li>043</li>
    <li>044</li>
    <li>045</li>
    <li>046</li>
    <li>047</li>
    <li>048</li>
    <li>049</li>
    <li>050</li>
    <li>051</li>
    <li>052</li>
    <li>053</li>
    <li>054</li>
    <li>055</li>
    <li>056</li>
    <li>057</li>
    <li>058</li>
    <li>059</li>
    <li>060</li>
    <li>061</li>
    <li>062</li>
    <li>063</li>
    <li>064</li>
    <li>065</li>
    <li>066</li>
    <li>067</li>
    <li>068</li>
    <li>069</li>
    <li>070</li>
    <li>071</li>
    <li>072</li>
    <li>073</li>
    <li>074</li>
    <li>075</li>
    <li>076</li>
    <li>077</li>
    <li>078</li>
    <li>079</li>
    <li>080</li>
    <li>081</li>
    <li>082</li>
    <li>083</li>
    <li>084</li>
    <li>085</li>
    <li>086</li>
    <li>087</li>
    <li>088</li>
    <li>089</li>
    <li>090</li>
    <li>091</li>
    <li>092</li>
    <li>093</li>
    <li>094</li>
    <li>095</li>
    <li>096</li>
    <li>097</li>
    <li>098</li>
    <li>099</li>
    <li>100</li>
</ul>
</body>
</html>
	`)
}
