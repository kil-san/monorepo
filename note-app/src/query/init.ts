import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from "@apollo/client";
import {onError} from '@apollo/link-error';

const uri = process.env.REACT_APP_GRAPHQL_DOMAIN

const initClient = () => {
  const httpLink = new HttpLink({
    uri,
    credentials: 'include'
  })

  const requestLink = new ApolloLink((operation, forward) => {
    return forward(operation);
  });

  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors) {
      graphQLErrors.forEach((error) => {
        console.log(error)
      })
    }
    if (networkError) {
      console.log(networkError)
    }
  })
  
  
  return new ApolloClient({
    link: ApolloLink.from([errorLink, requestLink, httpLink]),
    cache: new InMemoryCache({
      addTypename: true
    })
  })
}

export default initClient