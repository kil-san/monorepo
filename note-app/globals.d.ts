declare module '@apptheme' {
  import { ThemeOptions } from '@material-ui/core'

  export interface AppTheme extends ThemeOptions {
    mobile: string
    tablet: string
    maxContentWidth: string
    status?: {
      danger?: string
    },
    custom?: {
      black?: string
      footerGradient?: string
    }
  }
}
