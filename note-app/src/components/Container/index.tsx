import React from 'react'
import clsx from 'clsx'
import { makeStyles } from '@material-ui/styles'

interface IContainer {
  maxWidth?: boolean
  marginTop?: string
  marginBottom?: string
}

const useStyles = makeStyles(() => ({
  max_width: {
    maxWidth: '1244px',
    margin: 'auto'
  },
  margin: (props: IContainer) => ({
    marginTop: props.marginTop + '!important',
    marginBottom: props.marginBottom + '!important'
  })
}))

const Container: React.FC<IContainer> = ({
  maxWidth = true,
  children,
  marginTop = '7em',
  marginBottom = '7em'
}) => {
  const classes = useStyles({ marginTop, marginBottom })
  return (
    <div className={clsx(maxWidth && classes.max_width, classes.margin)}>
      {children}
    </div>
  )
}

export default Container
