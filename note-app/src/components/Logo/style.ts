import {
  makeStyles,
  createStyles,
} from '@material-ui/core/styles';

import { AppTheme } from '@apptheme'

const useStyles = makeStyles((theme: AppTheme) =>
  createStyles({
    logo_wrap: {
      width: '100%',
      margin: '0px',
      padding: '0px',
      display: 'grid',
      placeItems: 'center'
    }
  })
)

export default useStyles
