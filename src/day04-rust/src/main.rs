use std::fs::File;
use std::io::{BufReader, BufRead};
use crate::card::Card;

pub mod card;

fn main() -> std::io::Result<()> {
    // let file = File::open("../../inputs/day04/example.txt")?;
    let file = File::open("../../inputs/day04/input.txt")?;
    let reader = BufReader::new(file);
    let mut cards: Vec<Card> = vec![];
    let mut total_points = 0;
    let mut total_cards = 0;

    for line in reader.lines() {
        let line = line?;

        let mut card = create_card_from_line(line);
        card.check_winners();
        total_points += card.points.clone();

        cards.push(card);
    }

    let cloned_cards = cards.clone();

    for i in 0..cloned_cards.len() {
        let matching_cards: u32 = cloned_cards[i].winners.len() as u32;

        if matching_cards > 0 {
            for n in 1..(matching_cards + 1) {
                let index_of_copy = i + n as usize;

                if index_of_copy < cloned_cards.len() {
                    cards[index_of_copy].instances += cards[i].instances;
                }
            }
        }

        total_cards += cards[i].instances;

        println!("{:?}", cards[i]);
    }

    println!("{}", format!("{} {}", "Part One - Your cards are worth a total points of: ", total_points));
    println!("{}", format!("{} {}", "Part Two - Your amount of scratchcards iss: ", total_cards));

    Ok(())
}

fn create_card_from_line(line: String) -> Card {
    let split: Vec<_> = line.split_terminator(": ").collect();
    let id: u8 = split[0].strip_prefix("Card ").unwrap().replace(" ", "").parse::<u8>().unwrap();
    let sets: Vec<_> = split[1].split_terminator(" | ").collect();

    let card = Card {
        id: id,
        winning: sets[0].split_whitespace().collect::<Vec<&str>>().iter().map(|&s| s.parse().unwrap()).collect::<Vec<u8>>(),
        having: sets[1].split_whitespace().collect::<Vec<&str>>().iter().map(|&s| s.parse().unwrap()).collect::<Vec<u8>>(),
        winners: vec![],
        points: 0,
        instances: 1,
    };

    return card;
}
