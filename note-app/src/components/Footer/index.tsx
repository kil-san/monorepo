import React from 'react'
import useStyles from './style'
import { Typography } from '@material-ui/core'
import { Logo } from 'components'
import { Link } from 'react-router-dom'

const Footer = () => {
  const classes = useStyles()

  return (
    <div className={classes.bg}>
      <div className={classes.content_wrap}>
        <div className={classes.sub_container_1}>
          <div className={classes.logo_wrap}>
            <Logo className={classes.logo} iconColor='black' />
          </div>
          <Typography className={classes.copyright} paragraph>
            Copyright &copy; {new Date().getUTCFullYear()} Notte
          </Typography>
          <Typography className={classes.copyright} paragraph>
          support@notte.io
          </Typography>
          <Typography className={classes.content} paragraph>
            <Link className={classes.link} to='/terms'>
              Terms
            </Link>
          </Typography>
          <Typography className={classes.pipe_class} paragraph>
            {'  |  '}
          </Typography>
          <Typography className={classes.content} paragraph>
            <Link className={classes.link} to='/privacy'>
              Privacy
            </Link>
          </Typography>
        </div>
      </div>
    </div>
  )
}

export default Footer
