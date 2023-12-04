use std::fs::File;
use std::io::{BufReader, BufRead};
use std::cmp;

#[derive(Debug)]
struct Game {
    id: i32,
    red: i32,
    green: i32,
    blue: i32,
    possible: bool,
    power: i32,
}

struct Content {
    red: i32,
    green: i32,
    blue: i32,
}

fn main() -> std::io::Result<()> {
    // let file = File::open("../../inputs/day02/example-1.txt")?;
    let file = File::open("../../inputs/day02/input.txt")?;
    let reader = BufReader::new(file);

    let mut games: Vec<Game> = vec![];

    for line in reader.lines() {
        let line = line?;
        games.push(process_line(line));
    }

    let content = Content { red: 12, green: 13, blue: 14 };
    let mut sum_of_ids = 0;
    let mut sum_of_powers = 0;

    for mut game in &mut games {
        sum_of_powers += game.power;

        if game.red > content.red {
            game.possible = false;
            continue;
        };
        if game.green > content.green {
            game.possible = false;
            continue;
        };
        if game.blue > content.blue {
            game.possible = false;
            continue;
        };

        sum_of_ids += game.id;
    }


    println!("{}", sum_of_ids);
    println!("{}", sum_of_powers);

    // dbg!(games);
    Ok(())
}

fn process_line(line: String) -> Game {
    let split: Vec<_> = line.split_terminator(':').collect();
    let game_id: i32 = split[0].strip_prefix("Game ").unwrap().parse::<i32>().unwrap();
    let game_sets: Vec<_> = split[1].split_terminator(";").collect();

    let mut game = Game { id: game_id, red: 0, green: 0, blue: 0, possible: true, power: 0 };

    for game_set in game_sets.iter() {
        let cubes: Vec<_> = game_set.split_terminator(", ").collect();

        for cube in cubes.iter() {
            let cube_count: Vec<_> = cube.split_whitespace().collect();

            let count: i32 = cube_count[0].parse::<i32>().unwrap();

            match cube_count[1] {
                "red" => game.red = cmp::max(count, game.red),
                "green" => game.green = cmp::max(count, game.green),
                "blue" => game.blue = cmp::max(count, game.blue),
                _ => continue
            }
        }
    }

    game.power = game.red * game.green * game.blue;

    return game;
}


