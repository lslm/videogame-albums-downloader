# Videogame Soundtrack Downloader

### Introduction
Uma das coisas que eu mais gosto de fazer no meu tempo livre é ouvir as músicas dos meus jogos favoritos (o album de The Last Of Us é o meu favorito, diga-se de passagem). Mas nem sempre é possível encontrar esses albuns em plataformas de streaming como Spotify ou Deezer e particularmente eu não confio em muitos links da interwebs para encontrar essas músicas.

Entretanto, existe um site que acredito que seja o melhor lugar onde eu posso encontrar minhas músicas favoritas: o [Video Game Music](https://downloads.khinsider.com). Seu único problema é não permitir que seja baixado um album de uma vez, só sendo possível baixar uma música por vez. Isso na maioria das vezes se torna impraticável, além de ser perda de tempo.

É para isso que criei essa ferramenta em Go! Passando o link do album no site, esse programa usa técnicas de [web scrapping](https://en.wikipedia.org/wiki/Web_scraping) para baixar todas as músicas do album, poupando assim o tempo do ouvinte. 