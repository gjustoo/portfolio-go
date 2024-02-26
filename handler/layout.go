package handler

import (
	"github.com/gjustoo/portfolio-go/view/component"
	"github.com/gjustoo/portfolio-go/view/layout"
	"github.com/labstack/echo/v4"
)

type LayoutHandler struct{}

func (h LayoutHandler) HandleLayoutShow(c echo.Context) error {

	githubIcon := `<svg class="h-4 w-4 text-slate-200"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round">  <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22" /></svg>`

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
		component.NewTextComponent("RESUME", "<b>Work</b> <br/> <b> Backend Developer at Cleverpy (2021 - NOW)</b> <br/><span class=\"text-slate-400 mt-0 text-justify\">- Java, SpringBoot and PostrgeSQL (PostGIS).  </span> <br/> <span class=\"text-slate-400 text-justify\"> - Developed high performance middleware and API REST with complex algorithms. Obtaining data from IOT sensors and similar. </span><br/> <span class=\"text-slate-400 text-justify\"> - I also conducted the Backend technical interviews for the company </span> <br/><hr/><br/> <b class=\"text-slate-400 mb-12 text-justify\">Education</b> <br/> <b class=\"text-slate-400 mb-12 text-justify\">Software Development at LA VEREDA</b>  <span class=\"text-slate-400 text-justify\">- Graduated in 2021 </span>  <br/> <b class=\"text-slate-400 mb-12 text-justify\">Language</b>  <span class=\"text-slate-400 text-justify\"> Spanish (native), English (Highly Proficient) </span>"),
		component.NewTextComponent("CONTACT ME", "You can watch my github page  here <br/> GitHub "+githubIcon+", LinkedIn, Twitter, Instagram, Facebook, Email"),
	}

	return render(c, layout.Base(components))
}
