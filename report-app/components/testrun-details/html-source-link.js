import {FaCode as HtmlIcon} from 'react-icons/fa'

export default ({ reportId }) =>
  <a
    className="TestSourceLine-htmlSourceLink"
    href={`/html-source?id=${reportId}`}
    target="_blank"
  >
    <HtmlIcon />
    &nbsp;
    Inspect HTML ...

    <style jsx>{`
    .TestSourceLine-htmlSourceLink {
      font-size: 0.8em;
      cursor: pointer;
      font-family: roboto, noto;
    }
    `}</style>

  </a>
