#[derive(Clone, Debug,Default)]
pub struct Card {
    pub id: u8,
    pub winning: Vec<u8>,
    pub having: Vec<u8>,
    pub winners: Vec<u8>,
    pub points: u32,
    pub instances: u32,
}

impl Card {
    pub fn check_winners(&mut self) {
        let winners: Vec<u8> = self.having.iter().filter(|&&value| self.winning.contains(&value)).cloned().collect();
        let mut points: u32 = 0;

        if winners.len() > 0 {
            points = (2u32).pow((winners.len() - 1) as u32);
        }

        self.winners = winners.clone();
        self.points = points;
    }
}
