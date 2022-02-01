import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { ApolloProvider } from "@apollo/client";
import { Home, Note } from "pages";
import { NoteContextProvider } from "context";
import { initClient } from "query";
import { ThemeProvider } from "theme";

export default function App() {
  const graphQlClient = initClient();

  return (
    <ApolloProvider client={graphQlClient}>
      <ThemeProvider>
        <NoteContextProvider>
          <Router>
            <Switch>
              <Route exact path="/">
                <Home />
              </Route>
              <Route path="/note">
                <Note />
              </Route>
            </Switch>
          </Router>
        </NoteContextProvider>
      </ThemeProvider>
    </ApolloProvider>
  );
}
