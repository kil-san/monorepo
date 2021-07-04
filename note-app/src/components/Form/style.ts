import { makeStyles, createStyles } from '@material-ui/core/styles';
import { AppTheme } from '@apptheme'

const useStyles = makeStyles((theme: AppTheme) =>
  createStyles({
    root: {
      padding: '0em 1em'
    },
    title: {
      margin: '2em 0em'
    },
    hero: {
      width: '100%'
    },
    form: {
      height: '100%',
      padding: '1em'
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
    }
  })
)

export default useStyles
  