# Clima CLI

Clima CLI é uma ferramenta de linha de comando (CLI) simples para obter informações climáticas de uma cidade específica. Utilizando a API da OpenWeather, a Clima CLI fornece dados como temperatura atual, sensação térmica, temperatura máxima e mínima, e descrição das condições climáticas.

## Recursos

- Consulta informações climáticas de uma cidade especificada.
- Exibe os dados climáticos em formato de tabela no terminal.
- Utiliza variáveis de ambiente para armazenar de forma segura a chave de API da OpenWeather.

## Uso

1. Configure a variável de ambiente `OPENWEATHER_API_KEY` com sua chave de API da OpenWeather.
2. Execute a Clima CLI usando o comando `clima` seguido pelo nome da cidade desejada.

Exemplo:

```sh
export OPENWEATHER_API_KEY=SuaChaveDeAPIAqui
clima --city Natal
```

## Instalação

1. Certifique-se de ter o Go instalado em seu sistema.
2. Clone este repositório.
3. Navegue para o diretório do repositório.
4. Execute o comando `go build` para compilar o executável.
5. Mova o executável gerado para um local no seu `$PATH`.

## Contribuição

Contribuições são bem-vindas! Se você deseja melhorar este projeto, sinta-se à vontade para abrir problemas e enviar pull requests.

## Licença

Este projeto é licenciado sob a [Licença Pública Geral GNU versão 3 (GNU GPLv3)](LICENSE). Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.

---