use std::{env, fs, io, process::exit};

const MAX_RED: u32 = 12;
const MAX_GREEN: u32 = 13;
const MAX_BLUE: u32 = 14;

fn main() {
    let lines;
    let mut sum = 0;

    if env::args().len() == 2 {
        let filename = env::args().nth(1).unwrap();

        match fs::read_to_string(filename) {
            Ok(contents) => lines = contents,
            Err(_) => {
                println!("Could not open file");
                exit(1);
            }
        }
    } else {
        lines = io::read_to_string(io::stdin()).unwrap();
    }

    'game: for line in lines.lines() {
        let mut games = line.split(": ");

        let id: u32 = games
            .next()
            .expect("Wrong input data format")
            .replace("Game ", "")
            .parse()
            .expect("Wrong input data format");

        let rounds = games.next().expect("Wrong input data format").split("; ");
        for round in rounds {
            for cube in round.split(", ") {
                let v: Vec<&str> = cube.split(" ").collect();
                if v.len() != 2 {
                    println!("Wrong input data format");
                    exit(1);
                }

                let n: u32 = v[0].parse().expect("Wrong input data format");
                let color = v[1];

                match color {
                    "red" if MAX_RED < n => continue 'game,
                    "green" if MAX_GREEN < n => continue 'game,
                    "blue" if MAX_BLUE < n => continue 'game,
                    _ => continue
                }
            }
        }

        sum += id;
    }

    println!("{sum}");
}
