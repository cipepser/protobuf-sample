extern crate protobuf;

mod user;

use user::User;
use std::fs::File;
use std::io::{BufReader};
use protobuf::{CodedInputStream, Message};

fn main() {
    let file = File::open("./go_user.bin").expect("cannot open file");
    let mut bufferd_reader = BufReader::new(file);
    let mut cis = CodedInputStream::from_buffered_reader(&mut bufferd_reader);

    let mut u = User::new();
    u.merge_from(&mut cis).expect("fail to merge");

    println!("{}", u.get_name());
    println!("{}", u.get_age());
}
