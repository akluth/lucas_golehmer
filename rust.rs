// Copyright (c) 2018, 2019 by Alexander Kluth <deralex@cpan.org>
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.
extern crate ramp;
extern crate primal;
extern crate num_integer;

use ramp::Int;
use num_integer::Integer;

/// Check if given prime number is a Mersenne prime.
/// Returns true if it is a Mersenne prime and false
/// if the number is not a Mersenne prime, a prime at all
/// or even
/// 
/// # Example
/// 
/// ~~~
/// extern crate lucas_lehmer;
/// use lucas_lehmer::is_mersenne;
///
/// assert!(is_mersenne(19))
/// ~~~
pub fn is_mersenne(prime: usize) -> bool {
    if prime == 2 {
        return true;
    }

    if !primal::is_prime(prime as u64) || prime.is_even() {
        return false;
    }

    return check_mersenne(prime);
}

fn check_mersenne(prime: usize) -> bool {
    let mersenne: Int = Int::from(Int::from(2).pow(prime) - 1);
    let mut s: Int = Int::from(4);

    for _i in 2..prime {
        s = (&s * &s - 2) % &mersenne;
    }

    s == 0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_if_m19_is_prime() {
        assert!(check_mersenne(19 as usize));
    }
}