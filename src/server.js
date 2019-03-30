import express from "express";
import { ApolloServer, gql } from "apollo-server-express";
import axios from "axios";

const app = express();

const typeDefs = gql`
  type Query {
    users: [User]!
  }

  type User {
    name: String!
  }
`;

const resolvers = {
  Query: {
    users: async () => {
      const { data } = await axios.get("http://localhost:3000/api/users");
      return data;
    },
  },
};
const server = new ApolloServer({ typeDefs, resolvers });
server.applyMiddleware({ app });

export default app;
