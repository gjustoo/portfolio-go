package handler

import (
	"github.com/gjustoo/portfolio-go/view/component"
	"github.com/gjustoo/portfolio-go/view/layout"
	"github.com/labstack/echo/v4"
)

type LayoutHandler struct{}

func (h LayoutHandler) HandleLayoutShow(c echo.Context) error {

	projects := []*component.ProjectComponent{
		component.NewProjectComponent("My Portfolio", "Built using : Go, Echo, HTML and Tailwind.", "https://github.com/gjustoo/portfolio-go"),
		component.NewProjectComponent("Interpreter", "Built a simple interpreter in Go.", "https://github.com/gjustoo/monkey-interpreter"),
		component.NewProjectComponent("Facebook Scrapper", "Facebook marketplace scrapper/notifier, sends notification through Telegram when an ad satisfies the selected filters. </br> Built in Python with Selenium and  MongoDB.", "https://github.com/gjustoo/FM_scrapper/"),
		component.NewProjectComponent("Feed visualizer (Frontend)", "Built a custom Redit and Twitter feed visualizer using React and Redux.", "https://github.com/gjustoo/tfg-client"),
		component.NewProjectComponent("Fee visualizer (Backend)", "Built a feed visualizer API Rest and consumer using Java/SpringBoot and PostgreSQL.", "https://github.com/gjustoo/tfg-service"),
		component.NewProjectComponent("Token retrieval script", "Simple script util to retrieve a JWT from a \"login\" endpoint. Built int plain SHELL", "https://github.com/gjustoo/token_retrieval"),
	}

	components := []component.Component{
		component.NewTextComponent("ABOUT", "I'm a 23 year old backend developer living in Spain. Through persistence, self-discipline, and commitment, I achieved my goal of becoming a backend developer. What kept me on this journey is that I always find learning new things exciting and facing unfamiliar challenges. In addition to coding, I enjoy building side projects because it allows me to explore my creativity and accomplish exciting things."),
		component.NewTextComponent("SKILLS", "  <b>Backend</b>: Go, gin, Java, Spring, Python, MongoDB, SQL, PostgreSQL, PostGIS. <br/> <b>Frontend</b>: HTML, CSS, JavaScript, React, Redux.   <br/><b>Tools</b>: Git, GitHub, GitLab, Bitbucket, Jira. <br/>   <b>Other</b>: Agile, Scrum"),
		component.NewProjectLayout("PROJECTS", projects),
		component.NewTextComponent("RESUME", "Here's my exprecience <br/> My resume can be downloaded here"),
		component.NewTextComponent("MY LINKS", "Here are my links <br/> GitHub, LinkedIn, Twitter, Instagram, Facebook, Email"),
	}

	return render(c, layout.Base(components))
}
