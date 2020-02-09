interface Environment {
  port: number;
  pricing: string;
}

const env: Environment = {
  port: 3000,
  pricing: "pricing-api:50051"
};

export default env;
