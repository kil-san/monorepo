import React from 'react'
import useStyles from './style'
//import LogoImage from 'res/Logo/logo.svg.js'
import LogoImage from 'res/Logo/artey.png'

interface ILogo {
  className?: string
  onClick?: Function
  iconColor?: string
}

const Logo: React.FC<ILogo> = ({
  className,
  onClick = () => {},
  iconColor = '#15EB83'
}) => {
  const classes = useStyles()

  return (
    <div
      onClick={() => onClick()}
      className={`${classes.logo_wrap} ${className && className}`}
    >
      <img src={LogoImage} height={39} width={161}/>
    </div>
  )
}

export default Logo
