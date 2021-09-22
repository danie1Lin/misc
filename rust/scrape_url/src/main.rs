use std::fs;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let url = &args[1];
    let output = &args[2];

    let md = get_markdown(url.to_string());
    fs::write(&output, md.as_bytes()).unwrap();
    println!("coverted to markdown saved in {}.", output);
}

fn get_markdown(url: String) -> String {
    println!("Fetching url: {}", url);
    let body = reqwest::blocking::get(url).unwrap().text().unwrap();

    println!("Convert html to markdown...");
    html2md::parse_html(&body)
}

fn pi() -> f64 {
    3.1415926 // do not need semicolon
}

fn not_pi() {
    3.1415926; // return null(unit)
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(1 + 2, 4);
    }
}
