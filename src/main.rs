use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
pub struct Cli {
    #[command(subcommand)]
    pub command: Option<Command>,
}

#[derive(Subcommand, Debug, Clone)]
pub enum Command {
    Hello {
        #[arg(short, long, default_value = "world")]
        name: String,
    },
    Serve,
}

pub fn exec(name: String) {
    println!("hello {}!", name);
}

fn main() {
    let cli = Cli::parse();
    if let Some(v) = cli.command {
        match v {
            Command::Hello { name } => exec(name),
            Command::Serve => api::serve(),
        }
    }
}