const DECIMAL_PLACES = 100;

const round = (n: number): number =>
  Math.round((n + Number.EPSILON) * DECIMAL_PLACES) / DECIMAL_PLACES;

export default round;
