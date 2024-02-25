package handler

import (
	"github.com/gjustoo/portfolio-go/view/component"
	"github.com/gjustoo/portfolio-go/view/layout"
	"github.com/labstack/echo/v4"
)

type LayoutHandler struct{}

func (h LayoutHandler) HandleLayoutShow(c echo.Context) error {

	components := []component.Component{
		component.NewTextComponent("ABOUT", "I'm a 23 year old backend developer living in Spain. Through persistence, self-discipline, and commitment, I achieved my goal of becoming a backend developer. What kept me on this journey is that I always find learning new things exciting and facing unfamiliar challenges. In addition to coding, I enjoy building side projects because it allows me to explore my creativity and accomplish exciting things."),
		component.NewTextComponent("SKILLS", "  <b>Backend</b>: Go, gin, Java, Spring, Python, MongoDB, SQL, PostgreSQL, PostGIS. <br/> <b>Frontend</b>: HTML, CSS, JavaScript, React, Redux.   <br/><b>Tools</b>: Git, GitHub, GitLab, Bitbucket, Jira. <br/>   <b>Other</b>: Agile, Scrum"),
		component.NewTextComponent("PROJECTS", "PROJECTS !!!!"),
		component.NewTextComponent("RESUME", "Here's my exprecience <br/> My resume can be downloaded here"),
		component.NewTextComponent("MY LINKS", "Here are my links <br/> GitHub, LinkedIn, Twitter, Instagram, Facebook, Email"),
	}

	return render(c, layout.Base(components))
}
