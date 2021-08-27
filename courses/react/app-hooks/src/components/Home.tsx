export function Home() {
    return (
    <div>
        
        <div id="info-card1" className="flex flex-col items-center p-6 mt-24 md:flex-row md:justify-between
            md:px-20">
            <div className="info-text md:w-1/3">
                <h3>Some info</h3>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ultrices felis, pharetra vitae justo, mi ac enim pretium. Tortor arcu orci, eget proin ultricies ut. A, sit amet nibh pulvinar. Sed aliquet tellus mattis nascetur eu mauris purus.</p>
            </div>
            <div className="info-button">
                <a href="#info-card2">Next</a>
            </div>
        </div>

        <div id="info-card2" className="flex flex-col items-center p-6 md:flex-row-reverse
            md:justify-between md:px-20">
           <div className="info-text md:w-1/3 md:text-right">
                <h3>Other info</h3>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ultrices felis, pharetra vitae justo, mi ac enim pretium. Tortor arcu orci, eget proin ultricies ut. A, sit amet nibh pulvinar. Sed aliquet tellus mattis nascetur eu mauris purus.</p>
            </div>
            <div className="info-button">
                <span><a href="#subscribe">Next</a></span>
            </div>
        </div>

        <div id="subscribe" className="bg-yellow-600 p-4 md:p-10">
            <h2 className="font-sans text-3xl font-bold text-white my-2">Subscribe</h2>

            <div className="subscribe-content md:flex md:flex-row-reverse md:justify-between">
                <p className="font-serif text-base text-yellow-50 mb-10 md:w-5/12 md:text-right md:text-lg">Lorem ipsum dolor sit amet, consectetur adipiscing elit ut aliquam, purus sit amet luctus venenatis, lectus magna fringilla urna, porttitor rhoncus dolor purus non enim praesent elementum facilisis leo, vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam quis enim lobortis scelerisque fermentum dui faucibus in ornare quam viverra orci sagittis eu volutpat odio facilisis mauris sit amet massa vitae tortor condimentum lacinia quis vel eros donec ac odio</p>

                <form className="md:w-5/12 md:mt-4 md:space-y-2">
                    <input id="subscribe-email" type="email" name="email"
                        placeholder="john.doe@mail.com"
                        className="bg-yellow-400 text-black py-1 px-2 rounded w-full text-lg
                        placeholder-gray-500" />

                    <div className="subscribe-dropdown">
                        <div className="dropdown-main">
                            <span>Options</span>
                            <svg width="28" height="12" viewBox="0 0 28 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.945953 0.830811L13.662 13.6716C14.0636 14.0771 14.7224 14.0649 15.1087 13.6449L26.8919 0.830811" stroke="black"/>
                            </svg>
                        </div>
                        <div className="dropdown-selects hidden">
                            <span className="dropdown-selectable">Option 1</span>
                            <span className="dropdown-selectable">Option 2</span>
                        </div>
                    </div>

                    <button className="w-full bg-yellow-400 rounded py-2 font-sans font-bold text-lg" type="submit">Send</button>
                </form>
            </div>
        </div>

        <div className="h-28 bg-white"></div>
    </div>
    );
}

