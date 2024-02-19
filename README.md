<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">gomlb</h3>

  <p align="center">
    A terminal based application for viewing live and past MLB games and statistics
    <br />
    <br />
    <br />
    <a href="https://github.com/AxBolduc/gomlb">View Demo</a>
    ·
    <a href="https://github.com/AxBolduc/gomlb/issues">Report Bug</a>
    ·
    <a href="https://github.com/AxBolduc/gomlb/issues">Request Feature</a>

[![Stargazers][stars-shield]][stars-url]
[![GPL License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
        <a href="#installation">Installation</a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

![gomlb Screen Shot][product-demo]

`gomlb` is a terminal based user interface that allows you to scroll through both live and completed MLB games and browse scores, players, and statistics.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

- Go
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Bubble-table](https://github.com/Evertras/bubble-table)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Installation

```bash
   go install github.com/axbolduc/gomlb@latest
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

By default running the application will open the game list view with the games from today's date. If you want to specify a specific date to open up to you can do so by doing

```bash
gomlb -d YYYY-MM-DD
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [x] List of games for a given day
- [ ] Game View
  - [x] Score Display
  - [x] Line Score
  - [x] Batters Box Score
  - [x] Pitchers Box Score
  - [ ] Other Box Score Info
- [ ] Player View
  - [ ] Bio Information
  - [ ] Season Stats
  - [ ] Career Stats

See the [open issues](https://github.com/AxBolduc/gomlb/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the GPL License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Alex Bolduc - [Linkedin](https://linkedin.com/in/twitter_handle)

Project Link: [https://github.com/AxBolduc/gomlb](https://github.com/AxBolduc/gomlb)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [nbacli](https://github.com/dylantientcheu/nbacli/) for inspiration and code snippets

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[stars-shield]: https://img.shields.io/github/stars/AxBolduc/gomlb.svg?style=for-the-badge
[stars-url]: https://github.com/AxBolduc/gomlb/stargazers
[issues-shield]: https://img.shields.io/github/issues/AxBolduc/gomlb.svg?style=for-the-badge
[issues-url]: https://github.com/AxBolduc/gomlb/issues
[license-shield]: https://img.shields.io/github/license/AxBolduc/gomlb.svg?style=for-the-badge
[license-url]: https://github.com/AxBolduc/gomlb/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/axbolduc
[product-screenshot]: images/screenshot.png
[product-demo]: images/demo.gif
