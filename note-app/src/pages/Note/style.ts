import { makeStyles, createStyles } from '@material-ui/core/styles';
import { AppTheme } from '@apptheme'

const useStyles = makeStyles((theme: AppTheme) =>
  createStyles({
    root: {
      padding: '0em 1em'
    },
    title_container: {
      display: 'flex',
      margin: '2em 0em',
      flexDirection: 'row'
    },
    title_text: {
      flex: 'auto',
      textAlign: 'center'
    },
    back_button: {
      position: 'absolute'
    },
    hero: {
      width: '100%'
    },
    form: {
      height: '100%',
      padding: '1em',
      minHeight: '50vh'
    },
    form_text: {
      display: 'flex',
      margin: '1em 0em'
    },
    submit: {
      margin: '3em 0em'
    },
    theme_select: {
      display: 'block',
      marginTop: '2em',
      minWidth: 120
    },
    check_list: {
      display: 'flex',
      margin: '2em 0em'
    },
    body_text: {
      margin: '2em .5em'
    }
  })
)

export default useStyles
  