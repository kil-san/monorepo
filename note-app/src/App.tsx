import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { ApolloProvider } from "@apollo/client";
import { Home, Note } from "pages";
import { UseThemeContext } from "hooks";
import { initClient } from "query";
import { UseNoteContext } from "hooks";

export default function App() {
  const graphQlClient = initClient();

  return (
    <ApolloProvider client={graphQlClient}>
      <UseThemeContext>
        <UseNoteContext>
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
        </UseNoteContext>
      </UseThemeContext>
    </ApolloProvider>
  );
}
