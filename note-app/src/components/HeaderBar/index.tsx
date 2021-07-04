import React, { useState, useEffect } from 'react'
import useStyles from './style'
import {
  Button,
  IconButton,
  Menu,
  MenuItem
} from '@material-ui/core'
import { Menu as MenuIcon } from '@material-ui/icons'
import { Logo, Footer } from 'components'
import { Link } from 'react-router-dom'
import clsx from 'clsx'


interface IHeader {
  className?: string
  showFooter?: boolean
  subHeader?: any
}

const Header: React.FC<IHeader> = ({ className, subHeader, children, showFooter }) => {
  const [mobileMenu, setMobileMenu] = useState<any>(null)
  const classes = useStyles()

  const getStartedButton = (className: string, variant: any) => (
    <Link to='/signin'>
      <Button
        variant={variant}
        color='primary'
        className={className}
      >
        Explore
      </Button>
    </Link>
  )

  const buttons = (
    <>
      <Link to='/'>
        <Button className={classes.button}>Notes</Button>
      </Link>
    </>
  )

  const desktopButtons = (
    <>
      {buttons}
      {getStartedButton(classes.get_started_desktop, 'outlined')}
    </>
  )

  const mobileButtons = (
    <>
      <IconButton
        disableRipple
        onClick={e => setMobileMenu(e.target)}
        aria-label='Menu'
      >
        <MenuIcon className={classes.mobile_hamburger} />
      </IconButton>
      <Menu
        open={mobileMenu !== null}
        anchorEl={mobileMenu}
        onClose={() => setMobileMenu(null)}
      >
        <Link
          className={classes.icon_label}
          to='/'
          onClick={() => {
            setMobileMenu(null)
          }}
        >
          <MenuItem>Notes</MenuItem>
        </Link>
        {getStartedButton(classes.get_started_mobile, 'outlined')}
      </Menu>
    </>
  )

  return (
    <>
      <div className={classes.header_style}>
        <div className={classes.desktop_content_wrap}>
          <div className={clsx(classes.header_content, className)}>
            <Link to='/home'>
              <Logo />
            </Link>
            <div className={classes.desktop_buttons}>{desktopButtons}</div>
          </div>
        </div>
        <div className={classes.mobile_content_wrap}>
          <div className={clsx(classes.header_content, className)}>
            <Link to='/home'>
              <Logo className={classes.mobile_logo} />
            </Link>
            <div className={classes.mobile_buttons_wrap}>{mobileButtons}</div>
          </div>
        </div>
        
      </div>
      {subHeader && <div className={classes.subHeader}>{subHeader}</div>}
      <div className={classes.children}>{children}</div>
      {showFooter ? (
        <div className={classes.footer}>
          <Footer />
        </div>
      ) : (
        false
      )}
    </>
  )
}

export default Header
