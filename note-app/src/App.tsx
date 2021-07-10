import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { ApolloProvider } from "@apollo/client";
import { Home } from 'pages'
import { UseThemeContext } from 'hooks'
import { initClient } from 'query'

export default function App() {
  const graphQlClient = initClient()
  
  return (
    <ApolloProvider client={graphQlClient}>
      <Router>
        <Switch>
          <UseThemeContext>
            <Route path="/">
              <Home />
            </Route>
          </UseThemeContext>
        </Switch>
      </Router>
    </ApolloProvider>
  );
}
