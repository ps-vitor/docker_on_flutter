use reqwest;
use scraper::{Html, Selector};
use warp::Filter;

fn  scrap(url:  &str){
    let response=reqwest::blocking::get(url).unwrap();
    let html_content=response.text().unwrap();
    
    let document=Html::parse_document(&html_content);

    let selector=Selector::parse("div.content.clearfix").unwrap();

    if  let Some(div)=document.select(&selector).next(){
        println!("\n{}\n",  div.text().collect::<String>());
    }
}

#[tokio::main]
async   fn  runserver(){
    let route   =   warp::path("scrape")
        .and(warp::get())
        .map(||{
            warp::reply::json(&"Resultado do scraping")
        });

    warp::serve(route).run(([0,0,0,0],8080)).await;

}

fn main() {
    runserver();
    let url="https://telemedicina.paginas.ufsc.br/processo-seletivo/";
    scrap(url); 
}
