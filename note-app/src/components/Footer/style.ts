import { makeStyles, createStyles } from '@material-ui/core/styles';
import { AppTheme } from '@apptheme'

const useStyles = makeStyles((theme: AppTheme) =>
  createStyles({
    bg: {
      width: '100%',
      backgroundImage: theme.custom?.footerGradient
    },
    content_wrap: {
      bottom: '0',
      display: 'grid',
      gridTemplateColumns: 'repeat(4, auto)',
      width: '100%',
      margin: 'auto',
      maxWidth: theme.maxContentWidth,
      textAlign: 'center',
      zIndex: 2,
      paddingBottom: '2em'
    },
    content: {
      display: 'inline-block',
      color: theme.custom?.black
    },
    copyright: {
      display: 'block',
      width: 'fit-content',
      margin: '0px auto 0.5em auto',
      color: theme.custom?.black
    },
    link: {},
    logo: {
      display: 'inline-bock'
    },
    logo_wrap: {
      display: 'block',
      width: 'fit-content',
      margin: 'auto'
    },
    links_col_title: {
      margin: '0px 0px 8px',
      fontWeight: 'bold',
      textAlign: 'left',
      color: theme.custom?.black
    },
    sub_container_2: {
      display: 'block',
      margin: '1em',
      'padding-top': '40px'
    },
    sub_container_1: {
      'padding-top': '2em',
      margin: '1em',
      display: 'block'
    },
    pipe_class: {
      padding: '0px 5px 0px 5px',
      display: 'inline-block'
    },
    [theme.tablet]: {
      content_wrap: {
        gridTemplateColumns: '1fr auto 1fr',
        gridTemplateRows: 'repeat(2, auto)',
        width: 'fit-content',
        margin: 'auto'
      },
      sub_container_1: {
        gridColumn: '1/span 3',
        gridRow: 1
      },
      sub_container_2: {
        gridRow: 2
      }
    },
    [theme.mobile]: {
      content_wrap: {
        gridTemplateColumns: 'auto auto'
      },
      exchange_container: {
        display: 'none'
      },
      sub_container_1: {
        gridColumn: '1/span 2'
      },
      sub_container_2: {
        paddingTop: '0px'
      }
    }
  })
)

export default useStyles

