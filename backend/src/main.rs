use reqwest;
use scraper::{Html, Selector};
use warp::Filter;

fn scrap(url: &str) -> String {
    let response = reqwest::blocking::get(url).unwrap();
    let html_content = response.text().unwrap();
    let document = Html::parse_document(&html_content);
    let selector = Selector::parse("div.content.clearfix").unwrap();

    // Retornar diretamente o resultado, seja encontrado ou "Not found"
    if let Some(div) = document.select(&selector).next() {
        div.text().collect::<String>()
    } else {
        "Not found.".to_string()
    }
}

#[tokio::main]
async fn main() {
    let url_to_scrap = "https://telemedicina.paginas.ufsc.br/processo-seletivo/".to_string();

    let route = warp::path("scrape")
        .and(warp::get())
        .map(move || {
            let data = scrap(&url_to_scrap);
            warp::reply::json(&data)
        });

    println!("Server runs on http://0.0.0.0:8080/scrape");
    warp::serve(route).run(([0, 0, 0, 0], 8080)).await;
}
