use std::{env, fs, io, process::exit};

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

    for line in lines.lines() {
        let mut games = line.split(": ");
        games.next();

        let mut max_red = 0;
        let mut max_green = 0;
        let mut max_blue = 0;

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
                    "red" if n > max_red => max_red = n,
                    "green" if n > max_green => max_green = n,
                    "blue" if n > max_blue => max_blue = n,
                    _ => continue,
                }
            }
        }

        sum += max_red * max_green * max_blue;
    }

    println!("{sum}");
}
