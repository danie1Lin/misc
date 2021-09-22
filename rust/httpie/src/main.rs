use clap::{AppSettings, Clap};

#[derive(Clap, Debug)]
struct Opts {
    #[clap(short, long, parse(from_occurrences))]
    verbose: i32,
    
    #[clap(subcommand)]
    subcmd: SubCommand,
}

#[derive(Debug,Clap)]
enum SubCommand {
    Test(Test),
    DebugCmd(DebugCmd)
}

#[derive(Debug,Clap)]
struct Test {
    #[clap(short)]
    debug: bool
}

#[derive(Debug,Clap)]
struct DebugCmd {
    #[clap(short)]
    test: bool
}

fn main() {
    let opts = Opts::parse();
    println!("opts: {:?}", opts)
}
