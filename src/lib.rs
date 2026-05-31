use num_bigint::BigUint;
use num_traits::{One, Zero};

/// Checks whether `2^exponent - 1` is a Mersenne prime.
pub fn is_mersenne(exponent: usize) -> bool {
    if exponent == 2 {
        return true;
    }

    if exponent < 2 || !is_prime(exponent) {
        return false;
    }

    check_mersenne_prime(exponent)
}

fn is_prime(n: usize) -> bool {
    if n < 2 {
        return false;
    }

    if n == 2 {
        return true;
    }

    if n.is_multiple_of(2) {
        return false;
    }

    let mut divisor = 3usize;
    while divisor <= n / divisor {
        if n.is_multiple_of(divisor) {
            return false;
        }
        divisor += 2;
    }

    true
}

fn mersenne_number(exponent: usize) -> BigUint {
    (BigUint::one() << exponent) - BigUint::one()
}

fn check_mersenne_prime(exponent: usize) -> bool {
    let mersenne = mersenne_number(exponent);
    let mut s = BigUint::from(4u8);

    for _ in 0..exponent - 2 {
        s = ((&s * &s) - BigUint::from(2u8)) % &mersenne;
    }

    s.is_zero()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn detects_known_mersenne_prime_exponents() {
        for exponent in [2, 3, 5, 7, 13, 17, 19, 31] {
            assert!(is_mersenne(exponent), "{exponent}");
        }
    }

    #[test]
    fn rejects_known_non_mersenne_prime_exponents() {
        for exponent in [0, 1, 4, 11, 23, 29] {
            assert!(!is_mersenne(exponent), "{exponent}");
        }
    }

    #[test]
    fn builds_mersenne_number() {
        assert_eq!(mersenne_number(13).to_string(), "8191");
    }
}
