# Videogame Soundtrack Downloader

### Introdução
Uma das coisas que eu mais gosto de fazer no meu tempo livre é ouvir as músicas dos meus jogos favoritos (o album de The Last Of Us é o que eu mais ouço, diga-se de passagem). Mas nem sempre é possível encontrar esses albuns em plataformas de streaming como Spotify ou Deezer e particularmente eu não confio em muitos links da interwebs para encontrar essas músicas.

Entretanto, existe um site que acredito que seja o melhor lugar onde eu possa encontrar minhas músicas favoritas: o [Video Game Music](https://downloads.khinsider.com). Seu único problema é não permitir que os albums sejam baixados de uma só vez, só sendo possível baixar uma música por vez. Isso na maioria das vezes se torna impraticável, além de ser perda de tempo.

É nesse contexto que criei essa ferramenta em Go! Passando o link do album no site, esse programa usa técnicas de [web scraping](https://en.wikipedia.org/wiki/Web_scraping) para baixar todas as músicas do album, poupando assim o tempo do ouvinte.

### Executando
Antes é necessário compilar o projeto
```
$ GOOS=<os> GOARCH=<arch> go build -o videogame-soundtrack-downloader
```

Para executar, basta ir a linha de comando e executar o programa compilado, fornecendo os argumentos de URL do album e diretório de saída, como exemplificado:

```
$ videogame-soundtrack-downloader.exe https://downloads.khinsider.com/game-soundtracks/album/ace-combat-5-the-unsung-war-original-soundtrack C:\Users\lucas\Music\Ace Combat 5
```