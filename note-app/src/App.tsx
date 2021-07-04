import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { Home } from 'pages'
import { UseThemeContext } from 'hooks'

export default function App() {
  return (
    <Router>
      <Switch>
        <UseThemeContext>
          <Route path="/">
            <Home />
          </Route>
        </UseThemeContext>
      </Switch>
    </Router>
  );
}
