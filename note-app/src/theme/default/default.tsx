import { AppTheme } from '@apptheme'
import { createMuiTheme } from '@material-ui/core/styles';
import colors from './colors';
import config from './config';

const defaultTheme: AppTheme = {
  ...config,
  palette: {
    // Main Colors
    primary: { main: colors.primary },
    secondary: { main: colors.secondary }
  },
  custom: {
    ...colors.custom
  },
  overrides: {
    MuiTypography: {
      h1: {
        fontSize: '2.5rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      h2: {
        fontSize: '1.5rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      h3: {
        fontSize: '1.3rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      h4: {
        fontSize: '1.2rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      h5: {
        fontSize: '1.1rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      h6: {
        fontSize: '1rem',
        fontWeight: 500,
        color: colors.custom.black
      },
      subtitle1: {
        color: colors.custom.grey
      },
      body1: {
        color: colors.custom.grey
      },
      body2: {
        fontSize: '1rem'
      },
      caption: {
        color: colors.custom.grey
      }
    },
    MuiButton: {
      containedPrimary: {
        borderRadius: '3em',
        height: '3em',
        color: colors.custom.paper,
      }
    },
    MuiExpansionPanel: {
      root: {
        '&.Mui-expanded': {
          margin: '0px!important'
        }
      }
    },
    MuiPaper: {
      elevation1: {
        boxShadow: 'unset'
      }
    }
  }
}

const muiTheme = createMuiTheme(defaultTheme)

muiTheme.typography.h1 = {
  [muiTheme.breakpoints.down('xs')]: {
    fontSize: '1.5em'
  }
}

muiTheme.typography.h2 = {
  [muiTheme.breakpoints.down('xs')]: {
    fontSize: '1.25em',
    fontWeight: '700'
  }
}

muiTheme.typography.subtitle1 = {
  [muiTheme.breakpoints.down('xs')]: {
    fontSize: '0.8em'
  }
}

export default muiTheme
