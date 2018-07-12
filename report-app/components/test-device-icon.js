import DesktopIcon from 'react-icons/lib/fa/desktop'
import MobileIcon from 'react-icons/lib/fa/mobile'
import TabletIcon from 'react-icons/lib/fa/tablet'
import AndroidIcon from 'react-icons/lib/fa/android'
import IOSIcon from 'react-icons/lib/fa/apple'

export default ({deviceSettings}) =>
  <span>
  {
    deviceSettings && deviceSettings.Type === 'desktop' ? <DesktopIcon /> : null
  }
  {
    deviceSettings && deviceSettings.Type === 'mobile' ? <MobileIcon /> : null
  }
  {
    deviceSettings && deviceSettings.Type === 'tablet' ? <TabletIcon /> : null
  }
  {
    deviceSettings && deviceSettings.Type === 'android' ? <AndroidIcon /> : null
  }
  {
    deviceSettings && deviceSettings.Type === 'ios' ? <IOSIcon /> : null
  }
  </span>