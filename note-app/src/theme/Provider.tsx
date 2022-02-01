import React, { useState, useEffect} from 'react';
import { ThemeProvider, Theme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import { defaultTheme } from 'theme'

interface IUseThemeContext {
  theme?: Theme
}

const Provider: React.FC<IUseThemeContext> = ({ theme: appTheme, children }) => {
  const [theme, setTheme] = useState(defaultTheme)

  useEffect(() => {
    appTheme && setTheme(appTheme)
  }, [appTheme])

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline/>
      {children}
    </ThemeProvider>
  );
}

export default Provider
