import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { AppTheme } from '@apptheme'

type IProps = {
  drawerWidth: number
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
    },
    
    drawer: (props: IProps) => ({
      width: props.drawerWidth,
      flexShrink: 0,
    }),
    drawerPaper: props => ({
      width: props.drawerWidth,
    }),
    // necessary for content to be below app bar
    toolbar: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default
    }
  })
)

export default useStyles
