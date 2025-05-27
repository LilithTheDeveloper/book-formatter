-- This file allows to add custom divs to your Markdown and have them converted to sections within LaTex

function Div(el)
  if el.classes:includes("infobox") then
    return {
      pandoc.RawBlock("latex", "\\begin{infobox}"),
      el,
      pandoc.RawBlock("latex", "\\end{infobox}")
    }
  end
end